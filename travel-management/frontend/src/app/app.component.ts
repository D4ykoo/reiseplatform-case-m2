import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';
import { SearchBarComponent } from "./components/search-bar/search-bar.component";
import { ResultTableComponent } from "./components/result-table/result-table.component";
import { InfoCardComponent } from "./components/info-card/info-card.component";
import { EditPanelComponent } from "./components/edit-panel/edit-panel.component";
import { TravelOfferEditComponent } from "./components/travel-offer-edit/travel-offer-edit.component";
import { MainComponent } from "./components/main/main.component";

@Component({
    selector: 'app-root',
    standalone: true,
    templateUrl: './app.component.html',
    styleUrl: './app.component.css',
    imports: [CommonModule, RouterOutlet, NavBarComponent, SearchBarComponent, ResultTableComponent, InfoCardComponent, EditPanelComponent, TravelOfferEditComponent, MainComponent]
})
export class AppComponent {
  title = 'Travel';

  public onFileChange(event:any) {
    console.log(event);
    const file = event.files[0];
    const reader = new FileReader();
    reader.onload=()=> {
      const baseString64 = reader.result as string;
      console.log(baseString64);
    }

    if(file) {
      reader.readAsDataURL(file);
    }
  }
}
