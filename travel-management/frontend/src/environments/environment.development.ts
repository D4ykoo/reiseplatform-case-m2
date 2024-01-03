export const environment = {
    production: false,
    Hotel_API: window["env" as any]["API_URL_FR" as any] || "http://localhost:8086/api/v1/",
    debug: window["env" as any]["DEBUG_FR" as any] || false,
    Checkout_URL: window["env" as any]["CHECKOUT_URL" as any] || "http://localhost:8083/",
    Login_URL: window["env" as any]["LOGIN_URL" as any] || "http://localhost:8081/",
    Monitor_URL: window["env" as any]["MONITOR_URL" as any] || "http://localhost:8087/",

};
