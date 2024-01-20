export type AttestationType = 'direct' | 'indirect' | 'none';

export interface BeginLoginResponse {
    userId: string;
    challenge: string;
    replyingPartyId: string,
    userVerification: string;
    allowCredentials: Credential[]
    timeout: number;
}

export interface Credential {
    id: string;
    publicKey: string;
    attestationType: string;
    transport: string[];
}

export interface FinishLoginRequest {
    username: string;
    passkey: string;
    challenge: string;
}

export interface FinishLoginResponse {
    token: string;
}

export interface PublicKeyCredentialOptions {
    challenge: string;
    rp: PublicKeyRelyingPartyInfo;
    user: PublicKeyUserInfo;
    attestation: AttestationType;
    pubKeyCredParams: Array<PublicKeyCredParam>;
    timeout?: number;
    excludeCredentials?: string[];
    authenticatorSelection?: AuthenticatorSelection;
    extensions?: any;
    status?: string;
}

export interface AuthenticatorSelection {
    authenticatorAttachment?: 'platform' | 'cross-platform';
    requireResidentKey?: boolean;
    userVerification?: 'required' | 'preferred' | 'discouraged';
}

export interface PublicKeyRelyingPartyInfo {
    name: string;
    icon?: string;
    id?: string;
}

export interface PublicKeyUserInfo {
    id: string;
    name: string;
    displayName: string;
    icon?: string;
}

export interface PublicKeyCredParam {
    type: string;
    alg: number;
}