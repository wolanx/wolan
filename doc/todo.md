

wolan(pipeline)
	get config
	get code
	build
	push img
	deploy
	update ingress


- [ ] code(master)    -> webhook      -> wolan(pipeline)
- [ ] code(dev-a)     -> webhook      -> wolan(pipeline)
- [ ] code(dev-b)     -> webhook      -> wolan(pipeline)


/wolan
    /config
        /app-1
            /wolan.yaml
        /app-2
    /git
        /demo-a
        /demo-b


