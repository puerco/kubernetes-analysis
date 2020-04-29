FROM tensorflow/tensorflow:2.1.0-gpu-py3

RUN apt-get update \
    && apt-get install -y \
        gir1.2-gtk-3.0 \
        git \
        gobject-introspection \
        libgirepository1.0-dev \
        libcairo2-dev \
    && rm -rf /var/lib/apt/lists/*

RUN git config --global user.name "Kubeflow" && \
    git config --global user.email k8s@saschagrunert.de

COPY requirements.txt /

RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

ENTRYPOINT ["/bin/bash"]
