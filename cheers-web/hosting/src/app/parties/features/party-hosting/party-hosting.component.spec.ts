import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyHostingComponent } from './party-hosting.component';

describe('PartyHostingComponent', () => {
  let component: PartyHostingComponent;
  let fixture: ComponentFixture<PartyHostingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PartyHostingComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PartyHostingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
