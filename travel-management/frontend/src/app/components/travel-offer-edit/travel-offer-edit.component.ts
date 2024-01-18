import { CommonModule } from '@angular/common';
import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';
import { DropdownModule } from 'primeng/dropdown';
import { Hotel } from '../../models/hotel';
import { FormsModule } from '@angular/forms';
import { InputNumberModule } from 'primeng/inputnumber';
import { CalendarModule } from 'primeng/calendar';
import { EditorModule } from 'primeng/editor';
import { lastValueFrom } from 'rxjs';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { CreateTravel } from '../../models/requests';
import { Travel } from '../../models/travel';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { LoginService } from '../../services/login.service';
import { User } from '../../models/user';

@Component({
  selector: 'app-travel-offer-edit',
  standalone: true,
  imports: [
    DropdownModule,
    CommonModule,
    FormsModule,
    InputNumberModule,
    CalendarModule,
    EditorModule,
    ToastModule,
  ],
  templateUrl: './travel-offer-edit.component.html',
  styleUrl: './travel-offer-edit.component.css',
  providers: [MessageService],
})
export class TravelOfferEditComponent implements OnInit, OnChanges {
  @Input()
  editorMode!: string | undefined;

  user!: User;

  hotels!: Hotel[];
  selectedHotel: Hotel | undefined;
  travels!: Travel[];
  selectedTravel!: Travel | undefined;
  price!: number;
  public rangeDates: Date[] | undefined;
  description: string | undefined;

  constructor(
    private readonly httpClient: HttpClient,
    private messageService: MessageService,
    private loginService: LoginService
  ) { }

  ngOnInit(): void {
    this.loginService.user.subscribe((u) => this.user = u);
    this.setup();
  }

  public setup() {
    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'hotels')).then(
      (hotels) => {
        if (hotels) this.hotels = hotels as Hotel[];
      },
    );
  }
  submit() {
    let from = new Date();
    let to = new Date();

    if (this.rangeDates) {
      from = this.rangeDates[0];
      to = this.rangeDates[1];
      if (!to) {
        to = from;
      }
    }
    if (this.editorMode == 'New') {
      const createTravel: CreateTravel = {
        description: this.description as string,
        from: from.toISOString(),
        to: to.toISOString(),
        price: this.price,
        vendorid: this.user.id,
        vendorname: this.user.name,
      };
      lastValueFrom(
        this.httpClient.post(
          environment.Hotel_API +
          'hotels/' +
          this.selectedHotel?.id +
          '/travels',
          createTravel, { withCredentials: true }
        ),
      )
        .then((res) => {
          this.clear();
          if (res) {
            this.messageService.add({
              severity: 'success',
              summary: 'Success',
              detail:
                'The hotel offer has been created (' + (res as Travel).id + ')',
            });
          }
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
    if (this.editorMode == 'Edit') {
      const updateTravel: CreateTravel = {
        description: this.description as string,
        from: from.toISOString(),
        to: to.toISOString(),
        price: this.price,
        vendorid: this.user.id,
        vendorname: this.user.name,
      };
      lastValueFrom(
        this.httpClient.put(
          environment.Hotel_API +
          'hotels/' +
          this.selectedHotel?.id +
          '/travels/' +
          this.selectedTravel?.id,
          updateTravel, { withCredentials: true }
        ),
      )
        .then((res) => {
          if (res) {
            const tmp = res as Travel;
            this.messageService.add({
              severity: 'success',
              summary: 'Success',
              detail: 'Hotel offer has been updated',
            });
            if (this.selectedTravel) {
              this.selectedTravel.id = tmp.id;
              this.selectedTravel.createdat = tmp.createdat;
              this.selectedTravel.description = tmp.description;
              this.selectedTravel.from = tmp.from;
              this.selectedTravel.price = tmp.price;
              this.selectedTravel.to = tmp.to;
              this.selectedTravel.updatedat = tmp.updatedat;
              this.selectedTravel.vendorid = tmp.vendorid;
              this.selectedTravel.vendorname = tmp.vendorname;
              this.loadSettings();
            } else {
              this.selectedTravel = tmp;
            }
          }
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
  }

  loadSettings() {
    if (this.selectedTravel) {
      this.description = this.selectedTravel.description;
      this.price = this.selectedTravel.price;
      this.rangeDates = [
        new Date(this.selectedTravel.from),
        new Date(this.selectedTravel.from),
      ];
    }
  }

  loadTravels() {
    if (this.editorMode == 'Edit' && this.selectedHotel) {
      this.travels = this.selectedHotel.travels;
      this.selectedTravel = undefined;
    }
  }

  loadTraveldata() {
    if (
      this.editorMode == 'Edit' &&
      this.selectedHotel &&
      this.selectedTravel
    ) {
      this.description = this.selectedTravel.description;
      this.price = this.selectedTravel.price;
      this.rangeDates = [
        new Date(this.selectedTravel.from),
        new Date(this.selectedTravel.to),
      ];
    }
  }

  clear() {
    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'hotels')).then(
      (res) => {
        if (res) this.hotels = res as Hotel[];
      },
    );
    this.selectedHotel = undefined;
    this.description = '';
    this.price = 0;
    this.rangeDates = [];
    this.travels = [];
    this.selectedTravel = undefined;
  }

  ngOnChanges(changes: SimpleChanges): void {
    const mode = changes['editorMode'];
    if (
      mode.currentValue != mode.previousValue &&
      mode.currentValue == 'Edit'
    ) {
      this.setup();
    } else {
      this.clear();
    }
  }
  delete() {
    if (this.selectedHotel && this.selectedTravel) {
      lastValueFrom(
        this.httpClient.delete(
          environment.Hotel_API +
          'hotels/' +
          this.selectedHotel?.id +
          '/travels/' +
          this.selectedTravel.id, { withCredentials: true }
        ),
      )
        .then(() => {
          this.clear();
          this.messageService.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Hotel offer has been deleted',
          });
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
  }

  async handleAuthorizationError(err: HttpErrorResponse) {
    if (err.status == 401) {
      console.error('Invalid Authorization: ' + err.message);
      this.messageService.add({
        severity: 'error',
        summary: 'Invalid Authorization',
        detail: 'Redirect to login page',
      });
      await this.delay(3000);
      let url = environment.Login_URL as unknown as string;
      url = url + '?name=travmngt';
      window.open(url, '_self');
    } else {
      console.error('ERROR:', err);
    }
  }

  delay(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
}
