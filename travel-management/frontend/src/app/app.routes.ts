import { Routes } from '@angular/router';
import { MainComponent } from './components/main/main.component';
import { EditPanelComponent } from './components/edit-panel/edit-panel.component';

export const routes: Routes = [
  { path: '', component: MainComponent },
  { path: 'management', component: EditPanelComponent },
];
