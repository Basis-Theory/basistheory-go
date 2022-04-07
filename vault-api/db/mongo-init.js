// eslint-disable-next-line no-undef
db.createCollection('tokens');
db.tokens.createIndex({ tenant_id: 1, token_id: 1 }, { unique: false });
db.tokens.createIndex({ tenant_id: 1, token_id: 1, property_name: 1 }, { unique: true });

db.createCollection('applications');
db.applications.createIndex({ tenant_id: 1 }, { unique: false });
db.applications.createIndex({ app_id: 1 }, { unique: false });
db.applications.createIndex({ app_key: 1 }, { unique: false });
db.applications.createIndex({ tenant_id: 1, app_id: 1, app_key: 1 }, { unique: true });
// eslint-disable-next-line no-undef
