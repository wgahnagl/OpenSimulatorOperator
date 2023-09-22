# For Windows images, the base os version has to match the container os version
FROM mcr.microsoft.com/powershell:lts-windowsservercore-ltsc2022@sha256:9d08dc98c778d92da9d9ea480885ec113d9bc7c3c1fc2a389d93b89ec3f76369

SHELL ["pwsh", "-Command", "$ErrorActionPreference = 'Stop'; $ProgressPreference = 'SilentlyContinue';"]

WORKDIR /app
#EXPOSE 80

ARG BUILD_VERSION
ARG ARCHIVE_URL=https://github.com/vletoux/pingcastle/releases/download/

# download
#RUN wget -Uri https://github.com/vletoux/pingcastle/releases/download/2.10.1.0/PingCastle_2.10.1.0.zip -UseBasicParsing
#RUN powershell -Command Invoke-WebRequest -Method Get -Uri ${ARCHIVE_URL}${BUILD_VERSION}/PingCastle_${BUILD_VERSION}.zip -OutFile c:\pingcastle.zip
RUN pwsh -Command \
	Invoke-WebRequest -Method Get -Uri https://github.com/vletoux/pingcastle/releases/download/3.0.0.0/PingCastle_3.0.0.0.zip -OutFile c:\app\pingcastle.zip ; \
	Expand-Archive -Path c:\app\pingcastle.zip -DestinationPath c:\app ; \
	Remove-Item c:\app\pingcastle.zip -Force

#RUN dir
RUN dir c:\app

# validation
RUN c:\app\pingcastle.exe

#RUN Expand-Archive -Path c:\pingcastle.zip -DestinationPath c:\
#RUN Remove-Item c:\python-3.5.1.exe -Force

HEALTHCHECK NONE

CMD ["c:\app\pingcastle.exe"]
