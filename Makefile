TMP_DIR=tmp
COV_OUT=${TMP_DIR}/coverage.out
COV_HTML=${TMP_DIR}/coverage.html

.PHONY: test
test:
	[ -d ${TMP_DIR} ] || mkdir -p ${TMP_DIR}
	go test -coverprofile ${COV_OUT} && go tool cover -html=${COV_OUT} -o ${COV_HTML}

.PHONY: clean
clean:
	if [ -f ${COV_OUT} ]; then rm -v ${COV_OUT} ; fi
	if [ -f ${COV_HTML} ]; then rm -v ${COV_HTML} ; fi
