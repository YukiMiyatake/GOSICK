FROM ubuntu:16.04
LABEL maintener: YukiMiyatake<YukiMariaMiyatake@gmail.com>

WORKDIR /root
ENV HOME /root

#apt
RUN apt-get update && apt-get install -y curl unzip gzip software-properties-common --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

#node
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash \
    && apt update \
    && apt -y install gcc g++ make nodejs

RUN npm config set user 0 \
    && npm config set unsafe-perm true
RUN npm update -g

#RUN apt-get install -y libsecret-1-dev
#serverless framework
RUN npm install -g -y serverless serverless-azure-functions azure-functions-core-tools

#ngrok
RUN curl -O https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip \
    && unzip ngrok-stable-linux-amd64.zip
ENV PATH=$HOME/ngrok/bin:$PATH

#dotnet core
RUN curl -O https://packages.microsoft.com/config/ubuntu/18.04/packages-microsoft-prod.deb \
    && dpkg -i packages-microsoft-prod.deb \
    && add-apt-repository universe \
    && apt-get install -y apt-transport-https \
    && apt-get update \
    && apt-get install -y dotnet-sdk-2.2

#AWS-cli
RUN apt-get install -y python3-pip
RUN pip3 install awscli

WORKDIR /src

#serverless create --template azure-nodejs --path my-service --name my-unique-name
#dotnet resutore  dotnet build
#CMD [ngrok http -host-header="0.0.0.0:7071" 7071]


