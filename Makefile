.PHONY: help
help: 
	@echo "make master - build and flash master node"
	@echo "make slave - build and flash slave node"
	@echo "make help - show this message"


.PHONY: master
master:
	cd ./master && make all


.PHONY: slave
slave:
	cd ./slave && make all