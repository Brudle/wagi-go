#LOGDIR   := _scratch/logs
CACHEDIR := _scratch/cache

.PHONY: serve
serve:
	RUST_LOG=wagi=trace ./wagi -l '0.0.0.0:8080' -c modules.toml --module-cache ${CACHEDIR}

.PHONY: tail-logs
tail-logs:
	tail -f ${LOGDIR}/*/*
