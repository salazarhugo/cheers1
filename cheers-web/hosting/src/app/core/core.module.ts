import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CoreRoutingModule} from "./core-routing.module";
import {CoreComponent} from './feature/core.component';
import {SharedModule} from "../shared/shared.module";
import {MaterialModule} from "../material/material.module";
import {TopbarModule} from "../home/ui/topbar/topbar.module";
import { SidenavComponent } from './ui/sidenav/sidenav.component';
import {RouterModule} from "@angular/router";
import { SidenavItemComponent } from './ui/sidenav-item/sidenav-item.component';
import {UserShellModule} from "../users/feature/user-shell/user-shell.module";

@NgModule({
    declarations: [
        CoreComponent,
        SidenavComponent,
        SidenavItemComponent,
    ],
    imports: [
        CommonModule,
        CoreRoutingModule,
        RouterModule,
        SharedModule,
        MaterialModule,
        UserShellModule,
        TopbarModule,
    ],
    providers: [],
    exports: [
        SidenavComponent,
        SidenavItemComponent
    ]
})
export class CoreModule {
}
