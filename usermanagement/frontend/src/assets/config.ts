export const enviroment = {
  production: false,
  // apiUrl: "http://localhost:8082/api",
};

export const headerConf = {
  withCredentials: true,
};

export const APP_CONFIG = {
  apiUrl: import.meta.env.API_URL,
  checkoutUrl: import.meta.env.CHECKOUT_URL,
  travelManagagementUrl: import.meta.env.TRAVELMANAGEMENT_URL,
  monitoringUrl: import.meta.env.MONITORING_URL,
};

export default APP_CONFIG;
