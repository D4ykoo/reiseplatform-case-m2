import { Component } from '@angular/core';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { ResultTableComponent } from '../result-table/result-table.component';

@Component({
  selector: 'app-main',
  standalone: true,
  templateUrl: './main.component.html',
  styleUrl: './main.component.css',
  imports: [SearchBarComponent, ResultTableComponent],
})
export class MainComponent {}
