import { Component, OnDestroy, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';
import { SearchBarComponent } from './components/search-bar/search-bar.component';
import { ResultTableComponent } from './components/result-table/result-table.component';
import { InfoCardComponent } from './components/info-card/info-card.component';
import { EditPanelComponent } from './components/edit-panel/edit-panel.component';
import { TravelOfferEditComponent } from './components/travel-offer-edit/travel-offer-edit.component';
import { MainComponent } from './components/main/main.component';
import { LoginService } from './services/login.service';
import { User } from './models/user';

@Component({
  selector: 'app-root',
  standalone: true,
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  imports: [
    CommonModule,
    RouterOutlet,
    NavBarComponent,
    SearchBarComponent,
    ResultTableComponent,
    InfoCardComponent,
    EditPanelComponent,
    TravelOfferEditComponent,
    MainComponent,
  ],
})
export class AppComponent implements OnInit {

  constructor(private loginService: LoginService) {
  }
  ngOnInit(): void {
    this.loginService.getLoginStatus();
  }
  title = 'Travel';
}
