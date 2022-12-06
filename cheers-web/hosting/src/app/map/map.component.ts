import {Component, OnInit} from '@angular/core';
import * as mapboxgl from 'mapbox-gl';

@Component({
    selector: 'app-map',
    templateUrl: './map.component.html',
    styleUrls: ['./map.component.sass']
})
export class MapComponent implements OnInit {
    map!: mapboxgl.Map;
    style = 'mapbox://styles/mapbox/streets-v11';
    lat = 37.75;
    lng = -122.41;

    constructor() {
    }

    ngOnInit(): void {
    }

    onMapLoad(map: mapboxgl.Map) {
        map.resize()
    }
}
