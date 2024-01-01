export const environment = {
    production: false,
    HotelAPI: window["env" as any]["API_URL_FR" as any] || "http://localhost:8083/api/v1/",
    debug: window["env" as any]["DEBUG_FR" as any] || false
    
};
