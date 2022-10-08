
# 代理 Kubernetes API Server【Reverse Tunneling Dialer】

# 项目Tower概述
  - 背景：在多集群管理场景下，对于运行在私有网络、裸金属服务、VM、容器中等环境中的Kubernetes集群，纳管受限、无法提供统一的集群管理入口
    混合、多云环境中变得分散割裂、集群维护困难、无法对集群的全生命周期进行维护、应用交付受限

  - 应用场景：
    - 暴露本地集群到公网环境，本地Kubernetes集群上获取公共LoadBalancer
    - 受托管的集群可以运行在云上（ACK、CCE、TKE）或云下（IDC）、 self-hosted Kubernetes cluster
    - Kubernetes API服务器代理允许Kubernete集群之外的用户连接到可能无法访问的集群IP。
    - 允许访问仅在集群网络中公开的服务。
    - apiserver充当用户和集群内端点之间的代理和堡垒。
    
  - API Server
  - kube-apiserver 是 Kubernetes 最重要的核心组件之一，主要提供以下的功能
    - 提供集群管理的 REST API 接口，包括认证授权、数据校验以及集群状态变更等
    - 提供其他模块之间的数据交互和通信的枢纽（其他模块通过 API Server 查询或修改数据，只有 API Server 才直接操作 etcd）
  - kube-apiserver 提供了 Kubernetes 的 REST API，实现了认证、授权、准入控制等安全校验功能，同时也负责集群状态的存储操作（通过 etcd）。
  ![](kube-apiserver.png)
  - rest api:
    `
    curl -v -H 'Content-Type: application/json' \
    "http://localhost:${PORT}/api/v1/namespaces/default/pods/${POD}/status" >"${POD}-orig.json"
    `
  - Tunnel Server
    - 底层采用了HTTP进行传输，将TCP/UDP链接封装在HTTP隧道中，并且还使用了SSH对通信数据进行加密。
    - Architecture
    ![](chisel.png)
  - Tower组成核心模块
    - Tunnel Server
    - Proxy代理模块、充当Kubernetes API Server代理的服务端
      - k8sproxy.NewUpgradeAwareHandler
        ````
        func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        u := *req.URL
        u.Host = s.host
        u.Scheme = s.scheme

        fmt.Printf("req.URL:%s\nreq.Body:%s\nreq.Header:%s\nreq.Host:%s\n", req.URL, req.Body, req.Header, req.Host)

        if s.useBearerToken && len(s.bearerToken) > 0 {
        req = utilnet.CloneRequest(req)
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.bearerToken))
        }

        // we choose one httpClient randomly
        rand.Seed(time.Now().UnixNano())
        s.rwLock.RLock()
        index := rand.Intn(len(s.httpClient))
        klog.V(5).Infof("server %s current agent connection length %d, random slice index %d", s.name, len(s.httpClient), index)
        httpProxy := k8sproxy.NewUpgradeAwareHandler(&u, s.httpClient[index].Transport, false, false, s)
        s.rwLock.RUnlock()

        httpProxy.ServeHTTP(w, req)
        }
    - Agent模块应用在客户集群、对外暴露集群服务
       
      ````
      func (agent *Agent) connectionLoop() {
      var connectionErr error
      b := &backoff.Backoff{
      Factor: 1.4, // for faster reconnection
      Max:    agent.options.MaxRetryInterval,
      }
      for agent.running {
      if connectionErr != nil {
      attempt := int(b.Attempt())
      maxAttempt := agent.options.MaxRetryCount
      d := b.Duration()

    		msg := fmt.Sprintf("Connection error: %s", connectionErr)
    		if attempt > 0 {
    			msg += fmt.Sprintf(" (Attempt: %d", attempt)
    			if maxAttempt > 0 {
    				msg += fmt.Sprintf("/%d", maxAttempt)
    			}
    			msg += ")"
    		}
    		klog.Warning(msg)

    		if maxAttempt > 0 && attempt >= maxAttempt {
    			break
    		}
    		klog.Warningf("Retrying in %s...", d)
    		connectionErr = nil

    		sig := make(chan os.Signal, 1)
    		signal.Notify(sig, syscall.SIGHUP)
    		select {
    		case <-time.After(d):
    		case <-sig:
    		}
    		signal.Stop(sig)
    	}

    	dialer := websocket.Dialer{
    		ReadBufferSize:   1024,
    		WriteBufferSize:  1024,
    		HandshakeTimeout: 45 * time.Second,
    		Subprotocols:     []string{version.ProtocolVersion},
    		Proxy:            http.ProxyFromEnvironment,
    	}

    	wsHeaders := http.Header{}

    	wsConn, _, err := dialer.Dial(agent.options.Server, wsHeaders)
    	if err != nil {
    		connectionErr = err
    		continue
    	}

    	conn := utils.NewWebSocketConn(wsConn)
    	klog.V(4).Info("Handshaking...")
    	sshConn, chans, reqs, err := ssh.NewClientConn(conn, "", agent.sshConfig)
    	if err != nil {
    		if strings.Contains(err.Error(), "unable to authenticate") {
    			klog.Error("Authentication failed", err)
    		} else {
    			klog.Error(err)
    		}
    		break
    	}

    	conf, _ := agent.config.Marshal()
    	klog.V(4).Info("Sending config")
    	t0 := time.Now()
    	_, configErr, err := sshConn.SendRequest("config", true, conf)
    	if err != nil {
    		klog.Error("Config verification failed", err)
    		break
    	}

    	// we may encounter error like 'A session already allocated for this client.'
    	// continue can be helpful while we execute 'kubectl rollout restart -n kubesphere-system deployment cluster-agent'
    	// see issue #29
    	if len(configErr) > 0 {
    		connectionErr = errors.New(string(configErr))
    		continue
    	}

    	klog.V(2).Infof("Connected (Latency %s)", time.Since(t0))
    	b.Reset()
    	agent.sshConn = sshConn
    	go ssh.DiscardRequests(reqs)
    	go agent.connectStreams(chans)

    	err = sshConn.Wait()
    	agent.sshConn = nil
    	if err != nil && err != io.EOF {
    		connectionErr = err
    		continue
    	}
    	klog.V(2).Info("Disconnected")
        }
        close(agent.runningC)
        }  
  - Tower代理模块依赖关系
    - entity
      ![](remotedialer.png)
      - proxy服务端维护ws连接会话，用于代理转发客户端请求sessions会话map[string]*HTTPProxy
      - 基于raft算法实现的leaderelection选举机制
        ````
        lock, err := resourcelock.New(resourcelock.LeasesResourceLock,
				"kubesphere-system",
				"tower",
				kubernetesClient.CoreV1(),
				kubernetesClient.CoordinationV1(),
				resourcelock.ResourceLockConfig{
					Identity: id,
					EventRecorder: record.NewBroadcaster().NewRecorder(scheme.Scheme, v1.EventSource{
						Component: "tower",
					}),
				})

			if err != nil {
				klog.Fatalf("error creating lock: %v", err)
			}

			leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
				Lock:          lock,
				LeaseDuration: options.LeaderElection.LeaseDuration,
				RenewDeadline: options.LeaderElection.RenewDeadline,
				RetryPeriod:   options.LeaderElection.RetryPeriod,
				Callbacks: leaderelection.LeaderCallbacks{
					OnStartedLeading: run,
					OnStoppedLeading: func() {
						klog.Errorf("leadership lost")
						os.Exit(0)
           },
           },
           })
      - agent基于backoff的心跳维护机制、保证代理链接的可用性
    - Data flow
      ![](remotedialer-flow.png)
  - tower 应用
    - 启动 proxy
    `
      ./bin/proxy --ca-cert=/root/certs/ca.crt --ca-key=/root/certs/ca.key --v=5 --kubeconfig=/root/.kube/k3s.yaml
      I1008 19:29:33.540192    5357 options.go:46] CA set to "/root/certs/ca.crt".
      I1008 19:29:33.540445    5357 options.go:47] CA key file set to "/root/certs/ca.key".
      I1008 19:29:33.540473    5357 options.go:48] Host set to 0.0.0.0
      I1008 19:29:33.540490    5357 options.go:49] Agent port set to 8080.
      I1008 19:29:33.540503    5357 options.go:50] Kubeconfig set to "/root/.kube/k3s.yaml".
      I1008 19:29:33.540513    5357 options.go:51] Leader election set to false
      I1008 19:29:33.542759    5357 proxy.go:292] Listening on 0.0.0.0:8080...
      I1008 19:29:40.071618    5357 proxy.go:133] New agent connection
      I1008 19:29:40.071712    5357 proxy.go:141] Handshaking...
      I1008 19:29:40.222967    5357 proxy.go:343]  is connecting from 43.142.137.98:34642
      I1008 19:29:40.223129    5357 proxy.go:149] Verifying configuration
      I1008 19:29:40.253403    5357 proxy.go:198] IssueCertAndKey for [127.0.0.1]
      I1008 19:29:40.455966    5357 proxy_server.go:131] Proxy server local-test-kubernetes: starting http proxy on :6146, proxy address 127.0.0.1:6443
      I1008 19:29:40.456099    5357 proxy.go:238] Connection established with local-test
    `
    - 启动 agent
    `./bin/agent --name=local-test  --token=f6402697107e92f16457f00430e8e2dfb6580e22cea77e98f7eacd1a4458dbd7  --proxy-server=http://120.48.90.114:8080  --keepalive=10s  --kubernetes-service=kubernetes.default.svc:443  --v=5  --kubeconfig=/home/ubuntu/.kube/k3s.yaml
      I1008 19:30:42.701993   13589 agent.go:201] Handshaking...
      I1008 19:30:42.894385   13589 agent.go:213] Sending config
      I1008 19:30:43.027009   13589 agent.go:229] Connected (Latency 132.595984ms)
    `
    - 验证 获取proxy生成的代理集群配置文件、从cr中获取
      `
      kubectl get nodes
      NAME            STATUS   ROLES                  AGE    VERSION
      vm-4-5-ubuntu   Ready    control-plane,master   118d   v1.23.6+k3s1
      `
      - watch list pod .....
        `
        root@ls-3mx933WO:~/.kube#  kubectl get all -A
        NAMESPACE     NAME                                          READY   STATUS      RESTARTS      AGE
        kube-system   pod/helm-install-traefik-crd-cqsnk            0/1     Completed   0             118d
        kube-system   pod/helm-install-traefik-7gdgf                0/1     Completed   0             118d
        kube-system   pod/svclb-traefik-j5jq5                       2/2     Running     6 (29m ago)   118d
        kube-system   pod/traefik-df4ff85d6-95m2s                   1/1     Running     3 (29m ago)   118d
        kube-system   pod/coredns-d76bd69b-4tmwz                    1/1     Running     3 (29m ago)   118d
        kube-system   pod/local-path-provisioner-6c79684f77-skrtf   1/1     Running     6 (28m ago)   118d
        kube-system   pod/metrics-server-7cd5fcb6b7-j62tc           1/1     Running     4 (28m ago)   118d
        NAMESPACE     NAME                     TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
        default       service/kubernetes       ClusterIP      10.43.0.1      <none>        443/TCP                      118d
        kube-system   service/kube-dns         ClusterIP      10.43.0.10     <none>        53/UDP,53/TCP,9153/TCP       118d
        kube-system   service/metrics-server   ClusterIP      10.43.152.90   <none>        443/TCP                      118d
        kube-system   service/traefik          LoadBalancer   10.43.234.18   10.0.4.5      80:31519/TCP,443:30517/TCP   118d
        `