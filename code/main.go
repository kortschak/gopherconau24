sysman, err := sys.NewManager(
	rpc.NewKernel, device.NewManager, store, datadir, log, &level, addSource)
if err != nil {
	return internalError
}
defer sysman.Close()

changes := make(chan config.Change)
w, err := config.NewWatcher(ctx, cfgdir, changes, -1, log)
if err != nil {
	return internalError
}
go w.Watch(ctx)

cfgman := config.NewManager(log)
for cfg := range changes {
	if cfg.Err != nil {
		...
	}
	err = cfgman.Apply(cfg)
	if err != nil {
		...
	}
	unified, ..., err := cfgman.Unify(public.Schema)
	if err != nil {
		...
	}
	err = sysman.Configure(ctx, unified)
	if err != nil {
		...
	}
}
