export const environment = {
  production: true,
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  apiUrl: window['env']['apiUrl'] || 'http://localhost:8084/api/v1',
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  usermanagementUrl:
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
    window['env']['usermanagementUrl'] || 'http://localhost:8081/#',
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  hotelmanagementUrl:
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
    window['env']['hotelmanagementUrl'] || 'http://localhost:8085/travma',
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-expect-error
  monitoringUrl: window['env']['monitoringUrl'] || 'http://localhost:8087/*',
};
