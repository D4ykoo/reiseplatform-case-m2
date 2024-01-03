import { ApplicationConfig, importProvidersFrom } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { HttpClientModule } from '@angular/common/http';
import { HashLocationStrategy, LocationStrategy } from '@angular/common';

export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes), {provide: LocationStrategy, useClass: HashLocationStrategy}, importProvidersFrom(HttpClientModule)]
};
