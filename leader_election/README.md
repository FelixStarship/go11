# leader election实现选举 
在kubernetes的世界中，很多组件仅仅需要***一个实例***在运行，比如controller-manager或第三方的controller，但是为了***高可用性***，需要组件有***多个副本***，在发生故障的时候需要自动切换。因此，需要利用***leader election的机制多副本部署***，***单实例运行***的模式。

