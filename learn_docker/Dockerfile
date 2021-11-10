FROM python:3.8-alpine3.14

RUN mkdir -p /home/naga
RUN pip install pymongo

COPY . /home/naga

CMD ["python3", "/home/naga/learn_mongodb.py"]