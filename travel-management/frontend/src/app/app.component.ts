import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';
import { SearchBarComponent } from "./components/search-bar/search-bar.component";
import { ResultTableComponent } from "./components/result-table/result-table.component";
import { InfoCardComponent } from "./components/info-card/info-card.component";
import { FileUploadModule } from 'primeng/fileupload';

@Component({
    selector: 'app-root',
    standalone: true,
    templateUrl: './app.component.html',
    styleUrl: './app.component.css',
    imports: [CommonModule, RouterOutlet, NavBarComponent, SearchBarComponent, ResultTableComponent, InfoCardComponent, FileUploadModule]
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
