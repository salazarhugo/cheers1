import {Injectable} from '@angular/core';
import {ApiService} from "./api.service";
import {HttpClient} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {environment} from "../../../../environments/environment.prod";
import {CodeGeneratorResponse} from "google-protobuf/google/protobuf/compiler/plugin_pb";

@Injectable({
    providedIn: 'root'
})
export class MapboxGeocodingService {

    BASE_URL = "https://api.mapbox.com/geocoding/v5"

    constructor(
        private http: HttpClient,
    ) {
    }

    searchPlace(search_text: string): Observable<Feature[]> {
        return this.http.get<MapboxOutput>(`${this.BASE_URL}/mapbox.places/${search_text}.json?access_token=${environment.mapbox.accessToken}&language=en`)
            .pipe(map(res => res.features))
    }
}

export interface MapboxOutput {
    attribution: string
    features: Feature[]
    query: [];
}

export interface Feature {
    place_name: string
    geometry: Geometry
}

export interface Geometry {
    type: string
    coordinates: number[]
}
