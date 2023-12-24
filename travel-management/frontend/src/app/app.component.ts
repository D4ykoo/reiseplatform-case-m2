import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';
import { SearchBarComponent } from "./components/search-bar/search-bar.component";
import { ResultTableComponent } from "./components/result-table/result-table.component";
import { InfoCardComponent } from "./components/info-card/info-card.component";

@Component({
    selector: 'app-root',
    standalone: true,
    templateUrl: './app.component.html',
    styleUrl: './app.component.css',
    imports: [CommonModule, RouterOutlet, NavBarComponent, SearchBarComponent, ResultTableComponent, InfoCardComponent]
})
export class AppComponent {
  title = 'Travel';
}
