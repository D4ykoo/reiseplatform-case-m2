import { Tag } from './tag';
import { Picture } from './pictures';
import { Travel } from './travel';

export interface Hotel {
  id: number;
  hotelname: string;
  street: string;
  state: string;
  land: string;
  vendorid: number;
  vendorname: string;
  description: string;
  pictures: Picture[];
  tags: Tag[];
  travels: Travel[];
}
