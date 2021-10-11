CURRENT := $(abspath .)

dbschema:
	echo $(CURRENT)
	mysqldump -ufarm -pfarm123456 -d farm > $(CURRENT)/db/farm.sql;