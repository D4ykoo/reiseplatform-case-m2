import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { DropdownModule } from 'primeng/dropdown';
import { Hotel } from '../../models/hotel';
import { FormsModule } from '@angular/forms';
import { InputNumberModule } from 'primeng/inputnumber';
import { CalendarModule } from 'primeng/calendar';
import { EditorModule } from 'primeng/editor';

@Component({
  selector: 'app-travel-offer-edit',
  standalone: true,
  imports: [DropdownModule, CommonModule, FormsModule, InputNumberModule,CalendarModule,EditorModule],
  templateUrl: './travel-offer-edit.component.html',
  styleUrl: './travel-offer-edit.component.css'
})
export class TravelOfferEditComponent {
  hotels!: Hotel[] | undefined;
  selectedHotel: Hotel | undefined;
  price!: number;
  public rangeDates: Date[] | undefined;
  description: string | undefined;
}
