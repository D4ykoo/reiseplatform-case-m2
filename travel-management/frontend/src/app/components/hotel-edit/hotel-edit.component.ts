import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MultiSelectModule } from 'primeng/multiselect';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { EditorModule } from 'primeng/editor';
import { InputMaskModule } from 'primeng/inputmask';
import { FileUploadModule } from 'primeng/fileupload';
import { CommonModule } from '@angular/common';
import { ContextMenuModule } from 'primeng/contextmenu';
import { CreateHotel, CreatePicture } from '../../models/requests';
import { lastValueFrom } from 'rxjs';
import { Hotel } from '../../models/hotel';
import { Tag } from '../../models/tag';
import { DropdownModule } from 'primeng/dropdown';
import { Travel } from '../../models/travel';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { User } from '../../models/user';
import { LoginService } from '../../services/login.service';

@Component({
  selector: 'app-hotel-edit',
  standalone: true,
  imports: [
    MultiSelectModule,
    FormsModule,
    EditorModule,
    InputMaskModule,
    FileUploadModule,
    CommonModule,
    ContextMenuModule,
    DropdownModule,
    ToastModule,
  ],
  templateUrl: './hotel-edit.component.html',
  styleUrl: './hotel-edit.component.css',
  providers: [MessageService],
})
export class HotelEditComponent implements OnInit, OnChanges {
  @Input()
  editorMode!: string | undefined;

  hotels!: Hotel[];
  hotel!: Hotel | undefined;

  tags!: Tag[];
  selectedTags!: Tag[];
  description: string | undefined;
  /* eslint-disable  @typescript-eslint/no-explicit-any */
  images = new Array<any>();
  pictures = new Array<CreatePicture>();
  actions: MenuItem[] | undefined;
  hotelname!: string;
  street!: string;
  state!: string;
  land!: string;
  user!: User;

  constructor(
    private httpClient: HttpClient,
    private messageService: MessageService,
    private loginService: LoginService,
  ) {}
  /* eslint-disable  @typescript-eslint/no-explicit-any */
  onFileChange(event: any) {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];

      const reader = new FileReader();
      reader.onload = () => {
        this.images.push(reader.result);
        this.pictures.push({
          id: 0,
          description: '',
          payload: reader.result as string,
        });
      };

      reader.readAsDataURL(file);
    }
  }

  deletePic(index: number) {
    if (index > -1) {
      this.images.splice(index, 1);
      this.pictures.splice(index, 1);
    }
  }

  submit() {
    if (this.editorMode == 'New') {
      const createHotel: CreateHotel = {
        description: this.description as string,
        hotelname: this.hotelname,
        land: this.land,
        pictures: this.pictures,
        state: this.state,
        street: this.street,
        tagids: this.selectedTags,
        vendorid: this.user.id,
        vendorname: this.user.name,
      };
      lastValueFrom(
        this.httpClient.post(environment.Hotel_API + 'hotels', createHotel, {
          withCredentials: true,
        }),
      )
        .then((res) => {
          this.clear();
          if (res) {
            this.messageService.add({
              severity: 'success',
              summary: 'Success',
              detail:
                'The hotel offer has been created (' + (res as Hotel).id + ')',
            });
          }
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
    if (this.editorMode == 'Edit') {
      const UpdateHotel: Hotel = {
        description: this.description as string,
        hotelname: this.hotelname,
        land: this.land,
        pictures: this.pictures,
        state: this.state,
        street: this.street,
        tags: this.selectedTags,
        vendorid: this.user.id,
        vendorname: this.user.name,
        id: this.hotel?.id as number,
        travels: this.hotel?.travels as Travel[],
      };
      lastValueFrom(
        this.httpClient.put(
          environment.Hotel_API + 'hotels/' + this.hotel?.id,
          UpdateHotel,
          { withCredentials: true },
        ),
      )
        .then((res) => {
          if (res) {
            this.messageService.add({
              severity: 'success',
              summary: 'Success',
              detail: 'Hotel offer has been updated',
            });
            const tmp = res as Hotel;
            if (this.hotel) {
              this.hotel.description = tmp.description;
              this.hotel.hotelname = tmp.hotelname;
              this.hotel.id = tmp.id;
              this.hotel.land = tmp.land;
              this.hotel.pictures = tmp.pictures;
              this.hotel.state = tmp.state;
              this.hotel.street = tmp.street;
              this.hotel.tags = tmp.tags;
              this.hotel.vendorid = tmp.vendorid;
              this.hotel.vendorname = tmp.vendorname;
              this.hotel.travels = tmp.travels;
            } else {
              this.hotel = tmp;
            }
            this.loadSettings();
          }
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
  }

  clear() {
    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'tags')).then(
      (res) => {
        if (res) this.tags = res as Tag[];
      },
    );
    this.selectedTags = [];
    this.description = '';
    this.images = new Array<any>();
    this.pictures = new Array<CreatePicture>();
    this.hotelname = '';
    this.street = '';
    this.state = '';
    this.land = '';
  }

  setup() {
    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'tags')).then(
      (res) => {
        if (res) this.tags = res as Tag[];
      },
    );

    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'hotels')).then(
      (res) => {
        if (res) this.hotels = res as Hotel[];
      },
    );

    this.actions = [{ label: 'Delete', icon: 'pi pi-fw pi-trash' }];
  }

  loadSettings() {
    if (this.editorMode == 'Edit' && this.hotel) {
      this.selectedTags = this.hotel.tags;
      this.description = this.hotel.description;
      this.hotelname = this.hotel.hotelname;
      this.street = this.hotel.street;
      this.state = this.hotel.state;
      this.land = this.hotel.land;
      this.pictures = this.hotel.pictures;
      this.images = [];
      this.hotel.pictures.forEach((img) => this.images.push(img.payload));
    }
  }

  ngOnInit() {
    this.loginService.user.subscribe((u) => {
      this.user = u;
    });
    this.setup();
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
    lastValueFrom(
      this.httpClient.delete(
        environment.Hotel_API + 'hotels/' + (this.hotel?.id as number),
        { withCredentials: true },
      ),
    )
      .then(() => {
        this.messageService.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Hotel offer has been deleted',
        });
        this.clear();
        this.hotel = undefined;
        this.hotels = new Array<Hotel>();
        this.setup();
      })
      .catch((err) => {
        this.handleAuthorizationError(err);
      });
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
