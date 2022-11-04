import {ComponentFixture, TestBed} from '@angular/core/testing';

import {PartyTicketCreateComponent} from './party-ticket-create.component';


describe('PartyTicketCreateComponent', () => {
    let component: PartyTicketCreateComponent;
    let fixture: ComponentFixture<PartyTicketCreateComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [PartyTicketCreateComponent]
        })
            .compileComponents();

        fixture = TestBed.createComponent(PartyTicketCreateComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
