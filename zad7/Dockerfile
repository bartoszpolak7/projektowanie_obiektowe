FROM swift:6.3-jammy

RUN apt-get update && apt-get install -y \
    libsqlite3-dev \
    build-essential \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/vapor/toolbox.git /tmp/toolbox \
    && cd /tmp/toolbox \
    && make install \
    && rm -rf /tmp/toolbox
 
WORKDIR /workspaces
 
CMD ["/bin/bash"]