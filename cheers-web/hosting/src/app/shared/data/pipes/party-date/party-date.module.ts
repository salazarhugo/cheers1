import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PartyDatePipe } from './party-date.pipe';



@NgModule({
    declarations: [
        PartyDatePipe
    ],
    exports: [
        PartyDatePipe
    ],
    imports: [
        CommonModule
    ]
})
export class PartyDateModule { }
