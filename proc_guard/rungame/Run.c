#include "IInterface.h"
#include <Windows.h>
#include <Shellapi.h>
#include <stdio.h>

//函数实现
__declspec(dllexport) void __stdcall RunGame(wchar_t* path)
{
    //wchar_t message[256];
    //wchar_t format[] = L"Launching game: %ls";

    //swprintf_s(message, sizeof(message) / sizeof(wchar_t), format, path);

    //MessageBoxW(NULL, message, L"Message", MB_OK);
    ShellExecuteW(NULL, L"open", path, NULL, NULL, SW_SHOWNORMAL);
  //ShellExecuteW(NULL, L"open", path, NULL, path, SW_SHOWNORMAL); //具体路径和文件名,可以精确找到,不然只有文件名会闪退
}