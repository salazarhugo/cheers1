
gen:
	cd ./proto/ && dos2unix genproto.sh && dos2unix create-pb-descriptor.sh
	cd ./proto/ && ./genproto.sh && ./create-pb-descriptor.sh && echo "done"
