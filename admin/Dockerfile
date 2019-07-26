FROM python:3.7-slim

WORKDIR /usr/src/bananadm

COPY libbananadm libbananadm/
COPY bananadm setup.py ./

RUN pip install .

ENTRYPOINT ["/usr/local/bin/bananadm"]
CMD ["--help"]