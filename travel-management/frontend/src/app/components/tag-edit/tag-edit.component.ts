import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { Tag } from '../../models/tag';
import { User } from '../../models/user';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { LoginService } from '../../services/login.service';
import { MessageService } from 'primeng/api';
import { lastValueFrom } from 'rxjs';
import { environment } from '../../../environments/environment';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-tag-edit',
  standalone: true,
  imports: [
    ToastModule,
    DropdownModule,
    FormsModule,
    CommonModule,
    ButtonModule,
  ],
  templateUrl: './tag-edit.component.html',
  styleUrl: './tag-edit.component.css',
  providers: [MessageService],
})
export class TagEditComponent implements OnInit, OnChanges {
  @Input()
  editorMode!: string | undefined;

  tags!: Tag[];
  tag!: Tag | undefined;
  tagname!: string;
  user!: User;

  constructor(
    private httpClient: HttpClient,
    private messageService: MessageService,
    private loginService: LoginService,
  ) {}

  ngOnInit() {
    this.loginService.user.subscribe((u) => {
      this.user = u;
    });
    // this.loginService.checkLoginStatus();
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

  setup() {
    lastValueFrom(this.httpClient.get(environment.Hotel_API + 'tags')).then(
      (res) => {
        if (res) this.tags = res as Tag[];
      },
    );
  }

  submit() {
    if (this.editorMode == 'New') {
      const createTag: Tag = {
        id: 0,
        name: this.tagname,
      };
      lastValueFrom(
        this.httpClient.post(environment.Hotel_API + 'tags', createTag, {
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
                'The hotel offer has been created (' + (res as Tag).id + ')',
            });
          }
        })
        .catch((err) => this.handleAuthorizationError(err));
    }
    if (this.editorMode == 'Edit') {
      const UpdateHotel: Tag = {
        id: this.tag?.id as number,
        name: this.tagname,
      };
      lastValueFrom(
        this.httpClient.put(
          environment.Hotel_API + 'tags/' + this.tag?.id,
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
            const tmp = res as Tag;
            if (this.tag) {
              this.tag.id = tmp.id;
              this.tag.name = tmp.name;
            } else {
              this.tag = tmp;
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
    this.tagname = '';
  }

  loadSettings() {
    if (this.editorMode == 'Edit' && this.tag) {
      this.tagname = this.tag.name;
    }
  }

  delete() {
    lastValueFrom(
      this.httpClient.delete(
        environment.Hotel_API + 'tags/' + (this.tag?.id as number),
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
        this.tag = undefined;
        this.tags = new Array<Tag>();
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
