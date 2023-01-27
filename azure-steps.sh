# Check runtimes available
az webapp list-runtimes --os linux --output table

# Variables
APP_NAME="tour-of-heroes-api-go"
RESOURCE_GROUP="tour-of-heroes-golang"
LOCATION="westeurope"
SQL_SERVER_NAME="tour-of-heroes-sql"
DB_NAME="tour-of-heroes-db"

# Deploy API to Azure
az webapp up \
--runtime "GO:1.19" \
--name $APP_NAME \
--resource-group $RESOURCE_GROUP \
--sku B1 \
--os linux \
--location $LOCATION

# Create SQL Server
SQL_USER_NAME="tour-of-heroes"
SQL_PASSWORD='Password1!'

az sql server create \
--resource-group $RESOURCE_GROUP \
--name $SQL_SERVER_NAME \
--location $LOCATION \
--admin-user $SQL_USER_NAME \
--admin-password $SQL_PASSWORD

# Allow Azure App Service to access SQL Server
az sql server firewall-rule create \
--resource-group $RESOURCE_GROUP \
--server $SQL_SERVER_NAME \
--name "AllowAzureServices" \
--start-ip-address 0.0.0.0 --end-ip-address 0.0.0.0

# Create SQL Database
az sql db create \
--resource-group $RESOURCE_GROUP \
--server $SQL_SERVER_NAME \
--name $DB_NAME

# Add environment variable to Azure App Service
az webapp config appsettings set \
--resource-group $RESOURCE_GROUP \
--name $APP_NAME \
--settings "DB_CONNECTION_STRING=sqlserver://$SQL_USER_NAME:$SQL_PASSWORD@$SQL_SERVER_NAME.database.windows.net:1433?database=$DB_NAME"

# Browse to the Azure App Service
az webapp browse --resource-group $RESOURCE_GROUP --name $APP_NAME

# See logs
az webapp log tail --resource-group $RESOURCE_GROUP --name $APP_NAME
