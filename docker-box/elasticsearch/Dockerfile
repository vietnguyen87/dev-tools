FROM docker.elastic.co/elasticsearch/elasticsearch:7.13.2

# copy config folder
COPY config/analysis /usr/share/elasticsearch/config/analysis

RUN /usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-icu && \
    /usr/share/elasticsearch/bin/elasticsearch-plugin install x-pack && \
    /usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-phonetic

