FROM mono:latest

ENV VERSION=0.9.2.2

RUN curl http://dist.opensimulator.org/opensim-${VERSION}.tar.gz -s | tar -xz -C /home
WORKDIR /home
RUN mv opensim-${VERSION} opensim 
COPY Regions.ini /home/opensim/bin/Regions/Regions.ini
COPY OpenSim.ini /home/opensim/bin/OpenSim.ini
EXPOSE 9000
WORKDIR /opensim-0.9.0.0-rc1/bin

CMD [ "mono",  "./OpenSim.exe" ]
