import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import {HashLocationStrategy, Location, LocationStrategy} from '@angular/common';
import { provideAnimations } from '@angular/platform-browser/animations'
import { provideHttpClient } from "@angular/common/http";

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes), {provide: LocationStrategy, useClass: HashLocationStrategy},provideAnimations(),provideHttpClient()]
};
