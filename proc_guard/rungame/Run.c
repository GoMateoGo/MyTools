#include "IInterface.h"
#include <Windows.h>
#include <Shellapi.h>
#include <stdio.h>

//����ʵ��
__declspec(dllexport) void __stdcall RunGame(wchar_t* path)
{
    //wchar_t message[256];
    //wchar_t format[] = L"Launching game: %ls";

    //swprintf_s(message, sizeof(message) / sizeof(wchar_t), format, path);

    //MessageBoxW(NULL, message, L"Message", MB_OK);
    ShellExecuteW(NULL, L"open", path, NULL, NULL, SW_SHOWNORMAL);
  //ShellExecuteW(NULL, L"open", path, NULL, path, SW_SHOWNORMAL); //����·�����ļ���,���Ծ�ȷ�ҵ�,��Ȼֻ���ļ���������
}