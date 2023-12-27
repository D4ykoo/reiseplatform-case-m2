import { Tag } from "primeng/tag";

export interface CreateHotel {
    hotelname: string;
    street: string;
    state: string;
    land: string;
    vendorid: number;
    vendorname: string;
    description: string;
    pictures: CreatePicture[];
    tagids: Tag[];
}


export interface CreatePicture {
    id: number;
    description: string;
    payload: string;
}

export interface CreateTravel {
    vendorid: number;
    vendorname: string;
    from: string;
    to: string;
    price: number;
    description: string;
}