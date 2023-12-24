import { Component } from '@angular/core';
import { DataViewModule } from 'primeng/dataview';
import { FormsModule } from '@angular/forms';
import { TagModule } from 'primeng/tag';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-result-table',
  standalone: true,
  imports: [DataViewModule,FormsModule,TagModule,CommonModule],
  templateUrl: './result-table.component.html',
  styleUrl: './result-table.component.css'
})
export class ResultTableComponent {
  products!: string[];
}
