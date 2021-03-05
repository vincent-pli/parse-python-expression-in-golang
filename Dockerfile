FROM registry.access.redhat.com/ubi8/python-38 AS builder
WORKDIR /

USER root

RUN mkdir -p /parse
COPY _output/bin/parser.py /parse

RUN pip install --upgrade pip && \
    pip install --no-cache-dir git+https://github.com/danthedeckie/simpleeval.git@master && \
    pip install --no-cache-dir pyinstaller

RUN cd /parse && pyinstaller parser.py --onefile


FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /

COPY _output/bin/parse-expression /usr/local/bin

COPY --from=builder /parse/dist/parser /usr/local/bin

ENTRYPOINT []
CMD ["parse-expression"]
