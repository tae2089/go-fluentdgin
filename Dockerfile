FROM tae2089/ubuntu:1.1

# throw errors if Gemfile has been modified since Gemfile.lock
WORKDIR /home/ubuntu
USER ubuntu
COPY test.conf /home/ubuntu
RUN sudo apt-get update -y
RUN sudo apt-get install -y git vim patch make bzip2 autoconf automake libtool bison curl ruby-full
RUN sudo gem install bundler
RUN sudo gem install fluentd
RUN sudo gem install etc json oj webrick