# Create a network
docker network create sqlserver-vnet

# Create a container with Azure SQL edge
docker run \
--name azuresqledge \
--network sqlserver-vnet \
--cap-add SYS_PTRACE -e 'ACCEPT_EULA=1' \
-e 'MSSQL_SA_PASSWORD=Password1!' \
 -v mssqlserver_volume:/var/opt/mssql \
-p 1433:1433 \
-d mcr.microsoft.com/azure-sql-edge 

#Create a database
docker run -it --network sqlserver-vnet mcr.microsoft.com/mssql-tools
sqlcmd -S azuresqledge -U SA -P 'Password1!' -Q 'CREATE DATABASE [heroes]'
exit