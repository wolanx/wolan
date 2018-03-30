

- [ ] code(master)    -> webhook      -> wolan(pipeline)
- [ ] code(dev-a)     -> webhook      -> wolan(pipeline)
- [ ] code(dev-b)     -> webhook      -> wolan(pipeline)


wolan(pipeline)
	get config
	get code
	build
	push img
	deploy
	update ha
