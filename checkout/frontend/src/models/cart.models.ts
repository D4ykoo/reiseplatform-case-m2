export interface CombinedCart {
  cart: Cart | undefined;
  hotel: Hotel[] | undefined;
}

export interface Cart {
  id: number;
  user_id: number;
  paid: boolean;
  payment_method: string | undefined;
}

export interface Hotel {
  hotelname: string;
  land: string;
  vendorname: string;
  description: string;
  pictures: string;
  travels: Travel[];
}

export interface Travel {
  vendorname: string;
  from: string;
  to: string;
  price: number;
}
