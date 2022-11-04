import {TestBed} from '@angular/core/testing';

import {MapboxGeocodingService} from './mapbox-geocoding.service';

describe('MapboxGeocodingService', () => {
    let service: MapboxGeocodingService;

    beforeEach(() => {
        TestBed.configureTestingModule({});
        service = TestBed.inject(MapboxGeocodingService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
