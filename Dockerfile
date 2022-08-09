FROM centos:7
LABEL maintainer="734839030@qq.com"
ARG MODULE_NAME

WORKDIR /data/${MODULE_NAME}

COPY target /data/${MODULE_NAME}

RUN echo -e "#!/bin/bash\n \
./bin/${MODULE_NAME}" > bin/start.sh  \
    && chmod +x bin/start.sh

CMD ["sh","-c","./bin/start.sh"]