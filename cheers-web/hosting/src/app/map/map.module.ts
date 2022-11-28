import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {MapRoutingModule} from './map-routing.module';
import {MapComponent} from './map.component';
import {NgxMapboxGLModule} from "ngx-mapbox-gl";
import {environment} from "../../environments/environment";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [
        MapComponent,
    ],
    imports: [
        CommonModule,
        MapRoutingModule,
        NgxMapboxGLModule.withConfig({
            accessToken: environment.mapbox.accessToken,
        }),
        FlexModule
    ],
    exports: [
        MapComponent,
    ],
})
export class MapModule {
}
