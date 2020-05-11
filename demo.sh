
if [ -d "${EDUCATION_DIR}/log" ];
then
    echo "the file is already exists"
else
    mkdir -p ${EDUCATION_DIR}/log
fi
count=`ps -ef |grep wechat-service.jar |grep -v "grep" |wc -l`
echo $count
if [ 0 == $count ]
then
 echo "wechat-service  may not be running"
else
 echo "runing....."
 ps -ef | grep wechat-service.jar | grep -v grep | awk '{print $2}' | xargs kill -9
fi

nohup ${JAVA_HOME}/bin/java -Xms100M -Xmx100M -jar wechat-service.jar >> ./log/nohup`date +%Y-%m-%d`.out  2>&1 &