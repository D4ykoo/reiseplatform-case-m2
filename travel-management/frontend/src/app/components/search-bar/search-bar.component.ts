import { Component } from '@angular/core';
import { InputGroupModule } from 'primeng/inputgroup';
import { InputGroupAddonModule } from 'primeng/inputgroupaddon';
import { CalendarModule } from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import { MultiSelectModule } from 'primeng/multiselect';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { OfferService } from '../../services/offer.service';
import { Tag } from '../../models/tag';

@Component({
  selector: 'app-search-bar',
  standalone: true,
  imports: [InputGroupModule, InputGroupAddonModule, CalendarModule, FormsModule, MultiSelectModule],
  templateUrl: './search-bar.component.html',
  styleUrl: './search-bar.component.css'
})

export class SearchBarComponent {
  public rangeDates: Date[] | undefined;
  public tags!: Tag[];
  public selectedTags!: Tag[];
  public destination!: string;
  public hotelname!: string;

  constructor(private readonly httpClient: HttpClient, private readonly offerService: OfferService) {

  }

  ngOnInit() {
    this.httpClient.get(environment.Hotel_API + "tags").subscribe((fetchedTags) => {
      this.tags = (fetchedTags as Array<Tag>)
    })
  }

  public searchOffers() {
    let from;
    let to;

    if (this.rangeDates) {
      from = this.rangeDates[0];
      to = this.rangeDates[1];
      if (!to) {
        to = from
      }
    }
    this.offerService.fetchOffers(this.destination, this.hotelname, from, to, this.selectedTags);
  }
}
