FROM centos:7
RUN yum install libfreetype.so.6 -y yum install libpthread.so.0 -y yum install fontconfig -y 
COPY ["CLI/glasswallCLI"           , "/usr/bin"        ] 
COPY ["Lib/libglasswall.classic.so", "/usr/lib"        ]
COPY ["CLI/Configs/*"              , "/home/glasswall/"]
RUN ln -s /usr/lib/libglasswall.classic.so /usr/bin/libglasswall.classic.so
RUN ldconfig
RUN chmod +x /usr/bin/glasswallCLI

RUN yum install gettext -y

COPY Config.ini.tmpl /home/glasswall/
COPY entrypoint.sh /
RUN chmod 755 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
