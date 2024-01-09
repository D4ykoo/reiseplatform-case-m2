import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Tag } from 'primeng/tag';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class TagService {
  private tagSubject = new BehaviorSubject<Array<Tag>>(new Array<Tag>());
  public tags = this.tagSubject.asObservable();

  constructor(private readonly httpClient: HttpClient) {}
}
