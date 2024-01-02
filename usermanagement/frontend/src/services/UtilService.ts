export class UtilService {
    public getJwtFromToken() {
        const cookieName = "authTravel";
        const cookeValue = document.cookie.split('; ').find(cookie => {
            cookie.startsWith(`${cookieName}=`)
        })?.split('=')[1]
        return cookeValue
    }
}